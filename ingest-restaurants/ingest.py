import csv
import sqlite3
import sys

# Increase the CSV field size limit
csv.field_size_limit(sys.maxsize)

# Define the CSV file path
csv_file_path = 'zomato.csv'

# Connect to the SQLite database
conn = sqlite3.connect('../warehouse.db')
cursor = conn.cursor()

# Create the Restaurant table if it does not exist
create_table_query = '''
create table if not exists Restaurant (
    id integer primary key autoincrement,
    Name varchar(255) not null,
    Area varchar(255),
    Address varchar(2048),
    MapsLink varchar(512) unique,
    MapsRatingOutOf5 float
);
'''
cursor.execute(create_table_query)

# Function to process the rating and votes from the CSV
def process_rating(rate):
    try:
        return float(rate.split('/')[0])
    except:
        return None

# Read the CSV file and insert the data into the Restaurant table
with open(csv_file_path, mode='r', encoding='utf-8') as csvfile:
    csv_reader = csv.DictReader(csvfile)
    for row in csv_reader:
        name = row['name']
        area = row['location']
        address = row['address']
        maps_link = row['url']
        maps_rating_out_of_5 = process_rating(row['rate'])

        # Insert the data into the table
        insert_query = '''
        INSERT INTO Restaurant (Name, Area, Address, MapsLink, MapsRatingOutOf5)
        VALUES (?, ?, ?, ?, ?)
        ON CONFLICT(MapsLink) DO NOTHING;
        '''
        cursor.execute(insert_query, (name, area, address, maps_link, maps_rating_out_of_5))

        print(f"Added {name} - {area}")

# Commit the transaction and close the connection
conn.commit()
conn.close()
