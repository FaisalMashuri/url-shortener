'use client'
import React, { useState } from 'react'
import { Input } from './ui/input'
import { Button } from './ui/button'
import toast from 'react-hot-toast';

export default function ShortenForm() {
    const [url, setUrl] = useState<string>("");
    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            console.log(url)
            const response = await fetch('http://127.0.0.1:9999/api/v1/short-url', {
                method: "POST",
                headers: {"Content-Type" : "application/json"},
                body: JSON.stringify({
                    url
                })
            })
            await response.json()
            setUrl("")
            toast.success("Success create short url !!")
        } catch (error) {
            console.log("Error shortener url : ", error)
        }

    }
  return (
    <form onSubmit={handleSubmit} className='mb-4'>
        <div className='space-y-4'>
            <Input value={url} onChange={(e) => setUrl(e.target.value)} className='h-12' type="url" placeholder='Enter URL to shorten' required/>
            <Button className='w-full p-2'>Shorten Url</Button>
        </div>
    </form>
  )
}
