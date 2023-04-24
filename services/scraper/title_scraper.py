# https://developers.google.com/explorer-help/code-samples#python

from datetime import datetime
import os
from dotenv import load_dotenv

load_dotenv()
import time
from googleapiclient.discovery import build

scopes = ["https://www.googleapis.com/auth/youtube.force-ssl"]


# Filter out data that we don't need
def massageData(data):
    # Set shortcut
    s = data["snippet"]

    # Get root thumbnail URL
    baseUrl = s["thumbnails"]["default"]["url"]
    rootThumbnail = baseUrl.rstrip("default.jpg")

    # Format date
    stringDate = s["publishedAt"]
    dtobj = datetime.strptime(stringDate, "%Y-%m-%dT%H:%M:%S%z")

    # Construct and return object
    return {
        "videoId": data["id"]["videoId"],
        "title": s["title"],
        "description": s["description"],
        "publishedAt": dtobj,
        "channelId": s["channelId"],
        "channelTitle": s["channelTitle"],
        "rootThumbnail": rootThumbnail,
        "recordDate": None,
        "artists": [],
        "suggestedTags": [],
    }


def main():
    api_service_name = "youtube"
    api_version = "v3"
    developerKey = os.environ["YT_API_KEY"]

    youtube = build(api_service_name, api_version, developerKey=developerKey)

    # Load all artists and groups
    f = open("./groupsAndArtists.txt", "r")
    lines = f.readlines()
    f.close()
    line = lines[0]
    artists = line.split(",")

    for artist in artists:
        print(f"current {artist}")

        request = youtube.search().list(
            part="snippet",
            maxResults=10,
            q=f"{artist} fancam",
            type="video",
            videoEmbeddable="true",
        )
        response = request.execute()

        items = response["items"]

        for item in items:
            title = item["snippet"]["title"]
            f = open("./titles.txt", "a")
            f.write(title + "\n")
            f.close()

        print("sleeping for 5")
        time.sleep(5)


if __name__ == "__main__":
    main()
