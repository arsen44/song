import { useState, useEffect } from "react";
import { SONG_DELETE, SONGS, SONG_UPDATE } from "../../../helpers/constants";
import axios from "axios";

// Тип для песни
interface Song {
  id: number;
  album_id: string;
  song_name: string;
  text: string;
  release_date?: string;
  link?: string;
}

interface UpdateSongParams {
  updatedTitle: string;
  updatedGroup: string;
  updatedText: string;
  updatedReleaseDate?: string;
  updatedLink?: string;
}

interface ModalProps {
  song: Song;
  onClose: () => void;
  onSave: (
    updatedTitle: string,
    updatedGroup: string,
    updatedText: string,
    updatedReleaseDate?: string,
    updatedLink?: string
  ) => void;
}

export default function Song() {
  const [songs, setSongs] = useState<Song[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [message, setMessage] = useState<string | null>(null);

  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const [currentSong, setCurrentSong] = useState<Song | null>(null);

  useEffect(() => {
    axios
      .get(SONGS)
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

  const handleDeleteSong = async (id: number) => {
    try {
      const response = await axios.delete(SONG_DELETE(id), {
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (response.status === 200 || response.status === 204) {
        setMessage("Song deleted successfully!");
        setSongs((prevSongs) => prevSongs.filter((song) => song.id !== id));
      } else {
        throw new Error("Unexpected response status");
      }
    } catch (err: unknown) {
      console.error(err);
      const errorMessage = axios.isAxiosError(err) ? err.response?.data?.error || err.message : "Failed to delete song";
      setMessage(errorMessage);
    }
  };

  const handleUpdateSong = async (
    id: number,
    { updatedTitle, updatedGroup, updatedText, updatedReleaseDate, updatedLink }: UpdateSongParams
  ): Promise<void> => {
    try {
      const response = await axios.put(SONG_UPDATE(id), {
        song_name: updatedTitle,
        group_name: updatedGroup,
        text: updatedText,
        release_date: updatedReleaseDate,
        link: updatedLink,
      });

      if (response.status === 200) {
        const updatedSong: Song = response.data;

        setSongs((prevSongs) =>
          prevSongs.map((song) =>
            song.id === id
              ? {
                  ...song,
                  song_name: updatedSong.song_name,
                  group_name: updatedSong.group_name,
                  text: updatedSong.text,
                  release_date: updatedSong.release_date,
                  link: updatedSong.link,
                }
              : song
          )
        );

        setMessage("Song updated successfully!");
      } else {
        throw new Error("Failed to update song");
      }
    } catch (err) {
      console.error(err);
      setMessage("An error occurred while updating the song");
    }
  };

  const openModal = (song: Song) => {
    setCurrentSong(song);
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setCurrentSong(null);
  };

  return (
    <div className="container mx-auto p-6">
      <h1 className="text-3xl font-bold text-center text-blue-600 mb-6">Songs</h1>

      {loading ? (
        <p>Loading songs...</p>
      ) : (
        <ul className="space-y-4">
          {songs.map((song) => (
            <li
              key={song.id}
              className="p-4 border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-shadow"
            >
              <h2 className="text-lg font-semibold">{song.album.Title}</h2>
              <p className="text-gray-600">{song.song_name}</p>
              <p className="text-gray-600">{song.text}</p>
              <p className="text-gray-600">{song.link}</p>

              <div className="flex space-x-2 mt-2">
                {/* Update Song */}
                <button
                  onClick={() => openModal(song)}
                  className="bg-yellow-500 text-white px-4 py-2 rounded-md hover:bg-yellow-600"
                >
                  Обнвоить
                </button>

                {/* Delete Song */}
                <button
                  onClick={() => handleDeleteSong(song.id)}
                  className="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700"
                >
                  Удалить
                </button>
              </div>
            </li>
          ))}
        </ul>
      )}

      {message && (
        <p className={`mt-4 ${message.includes("successfully") ? "text-green-600" : "text-red-600"}`}>{message}</p>
      )}

      {isModalOpen && currentSong && (
        <Modal
          song={currentSong}
          onClose={closeModal}
          onSave={(updatedTitle, updatedGroup, updatedText, updatedReleaseDate, updatedLink) => {
            handleUpdateSong(currentSong.id, {
              updatedTitle,
              updatedGroup,
              updatedText,
              updatedReleaseDate,
              updatedLink,
            });
            closeModal();
          }}
        />
      )}
    </div>
  );
}

// Компонент модального окна
function Modal({ song, onClose, onSave }: ModalProps) {
  const [updatedTitle, setUpdatedTitle] = useState(song.song_name);
  const [updatedGroup, setUpdatedGroup] = useState(song.group_name);
  const [updatedText, setUpdatedText] = useState(song.text);
  const [updatedReleaseDate, setUpdatedReleaseDate] = useState(song.release_date || "");
  const [updatedLink, setUpdatedLink] = useState(song.link || "");

  const handleSubmit = () => {
    onSave(updatedTitle, updatedGroup, updatedText, updatedReleaseDate, updatedLink);
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
      <div className="bg-white p-6 rounded-lg shadow-lg w-96">
        <h2 className="text-2xl font-semibold mb-4">Изменить песню</h2>
        <div className="mb-4">
          <label className="block mb-2 font-medium">Назание</label>
          <input
            type="text"
            value={updatedGroup}
            onChange={(e) => setUpdatedGroup(e.target.value)}
            className="w-full border rounded-md px-3 py-2"
          />
        </div>
        <div className="mb-4">
          <label className="block mb-2 font-medium">Группа</label>
          <input
            type="text"
            value={updatedTitle}
            onChange={(e) => setUpdatedTitle(e.target.value)}
            className="w-full border rounded-md px-3 py-2"
          />
        </div>
        <div className="mb-4">
          <label className="block mb-2 font-medium">Текст</label>
          <textarea
            value={updatedText}
            onChange={(e) => setUpdatedText(e.target.value)}
            className="w-full border rounded-md px-3 py-2"
            rows={5}
          ></textarea>
        </div>
        <div className="mb-4">
          <label className="block mb-2 font-medium">Дата</label>
          <input
            type="date"
            value={updatedReleaseDate}
            onChange={(e) => setUpdatedReleaseDate(e.target.value)}
            className="w-full border rounded-md px-3 py-2"
          />
        </div>
        <div className="mb-4">
          <label className="block mb-2 font-medium">Song Link</label>
          <input
            type="url"
            value={updatedLink}
            onChange={(e) => setUpdatedLink(e.target.value)}
            className="w-full border rounded-md px-3 py-2"
          />
        </div>
        <div className="flex justify-end space-x-2">
          <button onClick={onClose} className="px-4 py-2 bg-gray-300 rounded-md">
            Выход
          </button>
          <button onClick={handleSubmit} className="px-4 py-2 bg-blue-600 text-white rounded-md">
            Сохранить
          </button>
        </div>
      </div>
    </div>
  );
}
