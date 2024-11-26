"use client";

import { useState } from "react";
import { ADD_SONG } from "../../helpers/constants";
import Song from "./component/song";

// Тип для песни
interface Song {
  id: number;
  group_name: string;
  song_name: string;
  text: string;
}

export default function Home() {
  const [songs, setSongs] = useState<Song[]>([]);
  const [title, setTitle] = useState<string>("");
  const [group, setGroup] = useState<string>("");
  const [message, setMessage] = useState<string | null>(null);

  // Обработчик добавления песни
  const handleAddSong = async (e: React.FormEvent) => {
    e.preventDefault();
    setMessage(null);

    try {
      const response = await fetch(ADD_SONG, {
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
        <h2 className="text-xl font-semibold mb-4">Добавить песню</h2>
        <div className="mb-4">
          <label className="block text-gray-700 mb-2" htmlFor="title">
            Группа:
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
            Песня:
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
          Добавить песню
        </button>
        {message && (
          <p className={`mt-4 ${message.includes("successfully") ? "text-green-600" : "text-red-600"}`}>{message}</p>
        )}
      </form>

      {/* Список песен */}
      <Song />
    </div>
  );
}
