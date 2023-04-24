import os
import grpc
import fancam_pb2
import fancam_pb2_grpc
from google.protobuf.timestamp_pb2 import Timestamp
from dotenv import load_dotenv
from datetime import datetime
from googleapiclient.discovery import build

load_dotenv()


def tstamp_from_dt(dt):
    """Takes in a datetime variable dt and returns a protobuf-compatible timestamp variable"""
    ts = Timestamp()
    ts.FromDatetime(dt)
    return ts


def massageData(data):
    """Takes in a JSON object from the list "items" returned by
    the search endpoint of YouTube Date API v3
    """
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
        "recordDate": dtobj,
        "suggestedTags": [],
    }


def scrapeTitles(artist):
    api_service_name = "youtube"
    api_version = "v3"
    developerKey = os.environ["YT_API_KEY"]

    youtube = build(api_service_name, api_version, developerKey=developerKey)

    request = youtube.search().list(
        part="snippet",
        maxResults=2,
        q=f"{artist} fancam",
        type="video",
        videoEmbeddable="true",
    )
    response = request.execute()

    items = response["items"]

    massagedData = list(map(massageData, items))

    return massagedData


def insertFancams(fancamList: dict, suggested_tags) -> None:
    print(fancamList)
    # Connects to Fancam microservice and creates fancams
    with grpc.insecure_channel("localhost:6001") as channel:
        stub = fancam_pb2_grpc.FancamStub(channel)
        fancams = []
        for document in fancamList:
            suggestedTags = fancam_pb2.SuggestedTags(
                en_artist=["set"],
                en_group=["sldf"],
                en_song=[";alksdf"],
                kr_artist=["l;askdjf"],
                kr_group=["laksdf"],
                kr_song=["sldfh"],
            )

            fancam = fancam_pb2.FancamObject(
                id=document["videoId"],
                title=document["title"],
                description=document["description"],
                publishedAt=tstamp_from_dt(document["publishedAt"]),
                channelId=document["channelId"],
                channelTitle=document["channelTitle"],
                rootThumbnail=document["rootThumbnail"],
                record_date=tstamp_from_dt(document["recordDate"]),
                suggested_tags=suggestedTags,
            )
            fancams.append(fancam)

        print(fancams)
        print("Attempting to insert fancams...")
        res = stub.CreateFancams(fancam_pb2.FancamList(fancams=fancams))
        print("Successfully inserted fancams")


def main():
    # Read file with all artists and groups
    with open("data/groupsAndArtists.txt") as f:
        line = f.readline()

        artists = line.split(",")

        for artist in artists:
            titles = scrapeTitles(artist)

            insertFancams(titles, "hi")


if __name__ == "__main__":
    main()
