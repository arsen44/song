const baseUrl = process.env.NEXT_PUBLIC_BASE_URL || "http://localhost:2152";

// Golang Microservice API
export const ADD_SONG = `${baseUrl}/add-song/`;
export const SONGS = `${baseUrl}/songs/`;
export const SONGS_CHANGE_TEXT = (id) => `${baseUrl}/songs/${id}/text/`;
export const SONG_UPDATE = (id) => `${baseUrl}/songs/${id}/`;
export const SONG_DELETE = (id) => `${baseUrl}/songs/${id}/`;
