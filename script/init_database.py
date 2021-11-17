from pymongo import MongoClient
from dotenv import load_dotenv
import os, json


load_dotenv()

# connect to MongoDB, change the << MONGODB URL >> to reflect your own connection string
HOST = os.getenv('MONGO_HOST')
PORT = os.getenv('MONGO_PORT')
DATABASE = os.getenv('MONGO_DATABASE')

uri = 'mongodb://{}:{}/{}'.format(HOST, PORT, DATABASE)
client = MongoClient( uri )
db = client[DATABASE]

PATH = 'script/dataset'
files = os.listdir(PATH)
files = set(filter(lambda f: f.endswith('.json'), files))
for f in files:
    filename = os.path.join(PATH, f)
    collection = f.split('.')[0]
    doc = json.loads(open(filename).read())
    
    db[collection].insert_many(doc)
