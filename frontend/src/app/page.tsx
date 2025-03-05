"use client";

import { useState } from "react";
import MusicPlayer from "./component/MusicPlayer/MusicPlayer";
import Song from "./component/song";

// Тип для песни
interface Song {
  id: number;
  group_name: string;
  song_name: string;
  text: string;
}

export default function Home() {
  return (
    <div className="container mx-auto p-6">
      <MusicPlayer />
      <Song />
    </div>
  );
}
