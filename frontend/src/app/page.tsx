"use client";

import { useState, useEffect } from "react";

// Тип для песни
interface Song {
  id: number;
  title: string;
  artist: string;
}

export default function Home() {
  const [songs, setSongs] = useState<Song[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [title, setTitle] = useState<string>("");
  const [group, setGroup] = useState<string>("");
  const [message, setMessage] = useState<string | null>(null);

  // Получение списка песен
  useEffect(() => {
    fetch("http://192.168.1.67:8080/songs/")
      .then((res) => res.json())
      .then((data: Song[]) => {
        setSongs(data);
        setLoading(false);
      })
      .catch((err) => console.log("Failed to fetch songs:", err));
  }, []);

  // Обработчик добавления песни
  const handleAddSong = async (e: React.FormEvent) => {
    e.preventDefault();
    setMessage(null);

    try {
      const response = await fetch("http://192.168.1.67:2152/add-song/", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ group: group, song: title }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to add song");
      }

      const newSong: Song = await response.json();
      setSongs((prevSongs) => [...prevSongs, newSong]); // Обновляем список песен
      setTitle("");
      setGroup("");
      setMessage("Song added successfully!");
    } catch (err) {
      console.error(err);
      setMessage((err as Error).message);
    }
  };

  return (
    <div className="container mx-auto p-6">
      <h1 className="text-3xl font-bold text-center text-blue-600 mb-6">Songs</h1>

      {/* Форма добавления песни */}
      <form onSubmit={handleAddSong} className="mb-6 p-4 border border-gray-200 rounded-lg shadow-sm">
        <h2 className="text-xl font-semibold mb-4">Add a Song</h2>
        <div className="mb-4">
          <label className="block text-gray-700 mb-2" htmlFor="title">
            Title:
          </label>
          <input
            id="title"
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
            className="w-full px-4 py-2 border border-gray-300 rounded-md"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 mb-2" htmlFor="artist">
            Artist:
          </label>
          <input
            id="artist"
            type="text"
            value={group}
            onChange={(e) => setGroup(e.target.value)}
            required
            className="w-full px-4 py-2 border border-gray-300 rounded-md"
          />
        </div>
        <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700">
          Add Song
        </button>
        {message && (
          <p className={`mt-4 ${message.includes("successfully") ? "text-green-600" : "text-red-600"}`}>{message}</p>
        )}
      </form>

      {/* Список песен */}
      <ul className="space-y-4">
        {songs.map((song) => (
          <li
            key={song.id}
            className="p-4 border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-shadow"
          >
            <h2 className="text-lg font-semibold">{song.title}</h2>
            <p className="text-gray-600">By {song.artist}</p>
          </li>
        ))}
      </ul>
    </div>
  );
}
