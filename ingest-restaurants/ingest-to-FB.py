import csv
import sys
import firebase_admin
from firebase_admin import credentials, firestore
import uuid

# Increase the CSV field size limit
csv.field_size_limit(sys.maxsize)

# Define the CSV file path
csv_file_path = 'last1k.csv'

# Initialize the Firebase Admin SDK with your service account key
cred = credentials.Certificate("/Users/smaran/Secrets/menumemory-4b957-firebase-adminsdk-pbiyp-177162e8d9.json")  # Replace with your key path
firebase_admin.initialize_app(cred)

# Create a reference to Firestore
db = firestore.client()

# Function to process the rating and votes from the CSV
def process_rating(rate):
    try:
        return float(rate.split('/')[0])
    except:
        return None

# Read the CSV file and insert the data into Firestore
with open(csv_file_path, mode='r', encoding='utf-8') as csvfile:
    csv_reader = csv.DictReader(csvfile)
    for row in csv_reader:
        name = row['name']
        area = row['location']
        address = row['address']
        maps_link = row['url']
        maps_rating_out_of_5 = process_rating(row['rate'])

        # Data to be stored in Firestore
        data = {
            'name': name,
            'area': area,
            'address': address,
            'maps_link': maps_link,
            'maps_rating_out_of_5': maps_rating_out_of_5
        }

        # Add a new document to the "Restaurant" collection, using the MapsLink as the unique identifier
        doc_ref = db.collection('restaurants').document(str(uuid.uuid1()))
        doc_ref.set(data, merge=True)

        print(f"Added {name} - {area}")

print("Data ingestion complete!")
