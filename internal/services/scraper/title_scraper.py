# https://developers.google.com/explorer-help/code-samples#python

import os
import time
from googleapiclient.discovery import build

scopes = ["https://www.googleapis.com/auth/youtube.force-ssl"]

# Filter out data that we don't need
def massageData(data):
    # Set shortcut
    s = data['snippet']

    # Get root thumbnail URL
    baseUrl = s['thumbnails']['default']['url']
    rootThumbnail = baseUrl.rstrip("default.jpg")

    # Format date
    stringDate = s['publishedAt']
    dtobj = datetime.strptime(stringDate, '%Y-%m-%dT%H:%M:%S%z')

    # Construct and return object
    return {
        "videoId": data['id']['videoId'],
        "publishedAt": dtobj,
        "channelId": s['channelId'],
        "channelTitle": s['channelTitle'],
        "title": s['title'],
        "description": s['description'],
        "rootThumbnail": rootThumbnail,
    }


def main():
    api_service_name = "youtube"
    api_version = "v3"
    developerKey = "AIzaSyDvr36bn93daf31wtSr-a3Gsw6Kpa0-_jc"

    youtube = build(api_service_name, api_version, developerKey=developerKey)

    f = open("./listArtists.txt", 'r')
    lines = f.readlines()
    f.close()
    line = lines[0]

    artists = line.split(",")

    for artist in artists:
        print(f'current {artist}')

        request = youtube.search().list(
            part="snippet",
            maxResults=10,
            q=f'{artist} fancam',
            type="video",
            videoEmbeddable="true"
        )
        response = request.execute()

        items = response['items']

        for item in items:
            title = item['snippet']['title']
            f = open("./titles.txt", 'a')
            f.write(title + "\n")
            f.close()
        print('sleeping for 5')
        time.sleep(5)
        
        


if __name__ == "__main__":
    main()
