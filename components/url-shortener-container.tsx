'use client'
import React, {useState} from 'react'
import ShortenForm from './shorten-form';
import UrlList from './url-list';

export default function UrlShortenerContainer() {
  const [refreshKey, setRefreshKey] = useState(0);
  const handleRefresh = () => {
    setRefreshKey((prev) => prev + 1)
  }
  return <div>
    <ShortenForm handleUrlShortned={handleRefresh} />
    <UrlList key={refreshKey} />
  </div>
}

