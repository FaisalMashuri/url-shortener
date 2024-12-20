"use client"
import { Toaster } from "react-hot-toast";
import UrlShortenerContainer from "./../components/url-shortener-container";
export default function Home() {
  return (
    <main className="mx-auto max-w-xl py-12 md:py-24 space-y-6">
      <Toaster
        toastOptions={{
          success: {
            style: {
              background: "green",
              color: "white"
            },
          },
          error: {
            style: {
              background: "red",
              color: "white"
            },
          },
        }}
      />
      <div className="space-y-2 text-center">
        <h1 className="text-3xl md:text-4xl font-bold">Url Shortner</h1>
        <p className="md:text-lg">Shortener your URLs and share then easily</p>
      </div>
      <UrlShortenerContainer />
    </main>
  );
}
