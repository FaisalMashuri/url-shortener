"use client";
import Link from "next/link";
import React, { useEffect, useState } from "react";
import { Button } from "./ui/button";
import { Check, CopyIcon, EyeIcon } from "lucide-react";

type Url = {
  id: number;
  originalUrl: string;
  shortUrl: string;
  view: number;
};

export default function UrlList() {
  const [urls, seturls] = useState<Url[]>([]);
  const [loading, setLoading] = useState(true);
  const [copiedUrls, setCopiedUrls] = useState<{ [key: string]: boolean }>({}); 

  console.log("urls : ", urls);

  const shortUrl = (url: string) => `${window.location.host}/${url}`;

  const getUrls = async () => {
    try {
      const response = await fetch("http://127.0.0.1:9999/api/v1/short-url");
      const data = await response.json();
      seturls(data.data); // Update state
    } catch (error) {
      console.log("Error get url : ", error);
    } finally {
      setLoading(false); // Pastikan loading selesai setelah fetch
    }
  };

  const handleCopyUrl = (code: string) => {
    // Set "copied" status untuk URL yang diklik
    const fullUrl = `${window.location.host}/${code}`
    navigator.clipboard.writeText(fullUrl).then(() => {
        setCopiedUrls((prev) => ({
            ...prev,
            [code]: true,
          }));
    })
   
    // Timeout untuk reset status "copied" setelah beberapa detik
    setTimeout(() => {
      setCopiedUrls((prev) => ({
        ...prev,
        [code]: false,
      }));
    }, 2000); // Reset setelah 2 detik
  };

  useEffect(() => {
    getUrls();
  }, [loading]);
  return (
    <div>
      <h2 className="text-2xl font-bold mb-2">Recent URLs</h2>
      <ul className="space-y-2">
        {loading ? (
          <p>Loading...</p> // Loading state
        ) : urls.length > 0 ? (
          urls.map((url) => (
            <li
              key={url.id}
              className="flex items-center gap-2 justify-between"
            >
              <Link
                href={`/${url.shortUrl}`}
                className="text-blue-500"
                target="_blank"
              >
                {shortUrl(url.shortUrl)}
              </Link>
              <div className="flex items-center gap-3">
                <Button
                  variant="ghost"
                  size="icon"
                  className="text-muted-foreground hover:bg-muted"
                  onClick={() => handleCopyUrl(url.shortUrl)}
                >
                  {copiedUrls[url.shortUrl] ? (
                    <Check className="w-4 h-4" />
                  ) : (
                    <CopyIcon className="w-4 h-4" />
                  )}

                  <span className="sr-only">Copy URL</span>
                </Button>
                <span className="flex items-center gap-2">
                  <EyeIcon className="h-4 w-4" />
                  {url.view} views
                </span>
              </div>
            </li>
          ))
        ) : (
          <p>No URLs found.</p> // Jika tidak ada data
        )}
      </ul>
    </div>
  );
}
