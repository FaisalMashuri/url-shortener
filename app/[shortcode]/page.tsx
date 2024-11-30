"use client";

import { useParams } from "next/navigation";
import { useRouter } from "next/router";  // Gunakan useRouter untuk mengambil parameter dari URL
import React, { useEffect, useState } from "react";

type Url = {
  id: number;
  shortUrl: string;
  originalUrl: string;
  view: number;
};

export default function RedirectPage() {
    const { shortcode } = useParams();  // Menggunakan router.query untuk mendapatkan parameter shortcode
  const [url, setUrl] = useState<Url | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const getUrl = async () => {
    if (!shortcode) return; // Jika shortcode tidak tersedia, jangan lakukan apa-apa

    try {
      const response = await fetch(
        `http://127.0.0.1:9999/api/v1/short-url/${shortcode}`
      );
      if (!response.ok) {
        throw new Error("Failed to fetch URL");
      }
      const { data }: { data: Url } = await response.json();
      setUrl(data);  // Simpan URL dalam state
    } catch (error: any) {
      setError(error.message || "An error occurred");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getUrl();  // Panggil getUrl ketika shortcode tersedia
  }, [shortcode]);  // Hanya memanggil getUrl ketika shortcode berubah

  if (loading) {
    return <p>Redirecting...</p>;
  }

  if (error) {
    return <p>Error: {error}</p>;
  }

  // Jika URL ditemukan, lakukan redirect
  if (url) {
    window.location.href = url.originalUrl;
    return null;
  }

  return <p>Redirecting...</p>;
}
