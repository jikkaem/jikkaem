import os
from dotenv import load_dotenv

load_dotenv()
from pymongo import MongoClient
from datetime import datetime

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


client = MongoClient("mongodb://localhost:10001/")


def insertIntoDb(data):
    db = client["jikkaem"]
    coll = db["fancams"]

    coll.insert_many(data)


def main():
    api_service_name = "youtube"
    api_version = "v3"
    developerKey = os.environ["YT_API_KEY"]

    youtube = build(api_service_name, api_version, developerKey=developerKey)

    focus = "Tzuyu"

    request = youtube.search().list(
        part="snippet",
        maxResults=50,
        q="tzuyu fancam",
        type="video",
        videoEmbeddable="true",
    )
    response = request.execute()

    items = response["items"]

    massagedData = list(map(massageData, items))

    insertIntoDb(massagedData)

    print("done")


if __name__ == "__main__":
    main()
