'use client'
import React, { useState } from 'react'
import { Input } from './ui/input'
import { Button } from './ui/button'
import toast from 'react-hot-toast';

interface ShortenFormProps {
    handleUrlShortned: () => void
}

export default function ShortenForm({handleUrlShortned} : ShortenFormProps) {
    const [url, setUrl] = useState<string>("");
    const [isLoading, setLoading] = useState<boolean>(false)
    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true)
        try {
            const response = await fetch('http://127.0.0.1:9999/api/v1/short-url', {
                method: "POST",
                headers: {"Content-Type" : "application/json"},
                body: JSON.stringify({
                    url
                })
            })
            await response.json()
            setUrl("")
            handleUrlShortned()
            toast.success("Success create short url !!")
            setLoading(false)
        } catch (error) {
            console.log("Error shortener url : ", error)
        }

    }
  return (
    <form onSubmit={handleSubmit} className='mb-4'>
        <div className='space-y-4'>
            <Input value={url} onChange={(e) => setUrl(e.target.value)} className='h-12' type="url" placeholder='Enter URL to shorten' required/>
            <Button className='w-full p-2' disabled={isLoading}>
                {isLoading ? "Shortnening ...." : "Short url"}
            </Button>
        </div>
    </form>
  )
}
