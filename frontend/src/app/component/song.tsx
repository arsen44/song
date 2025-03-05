import { useState, useEffect } from "react";
import { SONG_DELETE, GET_SONGS_All, SONG_UPDATE } from "../../../helpers/constants";
import axios from "axios";
import List from "./List/List";

export default function Song() {
  const [songs, setSongs] = useState<Song[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [message, setMessage] = useState<string | null>(null);

  useEffect(() => {
    axios
      .get(GET_SONGS_All)
      .then((response) => {
        const data = response.data;
        setSongs(data);
        console.log(data);
        setLoading(false);
      })
      .catch((error) => {
        console.error("Failed to fetch songs:", error);
        setLoading(false);
      });
  }, []);

  return <List data={songs} />;
}
