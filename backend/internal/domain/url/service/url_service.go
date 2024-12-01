package service

import (
	redis_client "backend/infrastructure/redis"
	"backend/internal/domain/url"
	"backend/internal/domain/url/dto/request"
	"backend/internal/domain/url/dto/response"
	"context"
	"encoding/json"
	"strconv"
	"time"
)

type urlService struct {
	repo      url.UrlRepository
	redisRepo redis_client.Repository
}

// GetUrlByShortUrl implements url.UrlService.
func (service *urlService) GetUrlByShortUrl(ctx context.Context, shortUrl string) (response.UrlResponse, error) {
	var res response.UrlResponse

	// Raw model untuk Redis (tanpa GORM fields)
	type dataModelRaw struct {
		ID          uint       `json:"ID"`                                                          // ID adalah uint
		CreatedAt   time.Time  `json:"CreatedAt" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // Waktu pembuatan
		UpdatedAt   time.Time  `json:"UpdatedAt" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // Waktu pembaruan
		DeletedAt   *time.Time `json:"DeletedAt" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // Nilai nullable, gunakan pointer
		OriginalUrl string     `json:"OriginalUrl"`                                                 // URL asli
		ShortUrl    string     `json:"ShortUrl"`                                                    // URL pendek
		View        int16      `json:"View"`                                                        // Jumlah tampilan
	}

	dataString, err := service.redisRepo.Get(ctx, shortUrl)
	if dataString != "" && err == nil {
		// Jika data terescaped, bersihkan
		cleanDataString, err := strconv.Unquote(dataString)
		if err != nil {
			return res, err
		}

		var dataModel dataModelRaw
		err = json.Unmarshal([]byte(cleanDataString), &dataModel)
		if err != nil {
			return res, err
		}

		var dataUrl url.Url
		dataUrl.ID = dataModel.ID
		dataUrl.OriginalUrl = dataModel.OriginalUrl
		dataUrl.ShortUrl = dataModel.ShortUrl
		dataUrl.View = dataModel.View
		dataUrl.AddViewUrl()

		// Simpan perubahan view ke database
		err = service.repo.UpdateViewShortUrl(ctx, dataUrl)
		if err != nil {
			return res, err
		}

		// Marshal data untuk Redis
		dataMarshal, err := json.Marshal(&dataUrl)
		if err != nil {
			return res, err
		}

		// Simpan data ke Redis dengan durasi 2 jam
		expiration := time.Hour * 2
		_ = service.redisRepo.Save(ctx, dataUrl.ShortUrl, string(dataMarshal), expiration)

		res = dataUrl.ToResponse()
		return res, nil
	}

	// Jika data tidak ditemukan di Redis, ambil dari database
	data, errGetData := service.repo.GetUrlByShortUrl(ctx, shortUrl)
	if errGetData != nil {
		return res, errGetData
	}

	// Simpan ke Redis
	dataMarshal, err := json.Marshal(&data)
	if err != nil {
		return res, err
	}

	// Atur durasi expired Redis
	twoHourExpired := time.Hour * 2
	_ = service.redisRepo.Save(ctx, data.ShortUrl, string(dataMarshal), twoHourExpired)

	// Tambah view
	data.AddViewUrl()
	if err := service.repo.UpdateViewShortUrl(ctx, data); err != nil {
		return res, err
	}

	// Konversi ke response
	res = data.ToResponse()
	return res, nil
}

// GetAllUrl implements url.UrlService.
func (service *urlService) GetAllUrl(ctx context.Context) ([]response.UrlResponse, error) {
	var response []response.UrlResponse
	data, err := service.repo.GetAllUrl(ctx)
	if err != nil {
		return response, err
	}
	for _, url := range data {
		response = append(response, url.ToResponse())
	}
	return response, nil
}

// CreateShortUrl implements url.UrlService.
func (service *urlService) CreateShortUrl(ctx context.Context, data request.UrlRequest) (string, error) {
	dataUrl, err := url.NewUser(data)
	if err != nil {
		return "", err
	}
	errCreateUrl := service.repo.CreateShortUrl(ctx, &dataUrl)
	if errCreateUrl != nil {
		return "", errCreateUrl
	}
	return dataUrl.ShortUrl, nil

}

func NewUrlService(repo url.UrlRepository, redisRepo redis_client.Repository) url.UrlService {
	return &urlService{
		repo:      repo,
		redisRepo: redisRepo,
	}
}
