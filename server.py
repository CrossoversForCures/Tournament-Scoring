from flask import Flask, redirect, url_for, render_template, jsonify
from pymongo import MongoClient

app = Flask(__name__)

client = MongoClient('mongodb://localhost:27017/')
db = client['E4ETeams2024']
collection = db['users']

@app.route("/")
def welcome():
   return render_template('admin.html')

def add_team():
    player = {
       'name': "test",
       'division': "5/6Boys"
    }
    collection.insert_one(player)
    return jsonify({'message': 'User added successfully'})

if __name__ == "__main__":
    app.run(debug=True)