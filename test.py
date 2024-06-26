from pymongo import MongoClient

def test_connection():
    client = MongoClient('mongodb://localhost:27017/')
    db = client['E4ETeams2024']
    collection = db['5/6Boys']

    player = {
        'name': "test",
        'game_one_win': True,
        'game_one_score': 25,
        'game_two_win': False,
        'game_two_score': 15
        }

    insert_result = collection.insert_one(player)
    print(f"Inserted document ID: {insert_result.inserted_id}")

    # Retrieve the document
    retrieved_document = collection.find_one({'name': 'test'})
    print(f"Retrieved document: {retrieved_document}")

    # Close the connection
    client.close()

if __name__ == "__main__":
    test_connection()


