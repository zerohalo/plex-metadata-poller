# plex-metadata-poller

Usage:
PLEX_TOKEN="xxxxxxxxxxx" PLEX_SERVER="xxxxxxxxxx.plex.direct:32400" PLEX_CLIENT_NAME="Plexamp" PLEX_USER_NAME="zerohalo" plex-metadata-poller

Writes a now_playing.txt to the current directory with the format:

```
Title: Tepid Bile
Artist: Tipper
Album: Tip Hop
```

This was created for use with Rogue Amoeba's Audio Hijack software to insert metadata into Icecast streams.
