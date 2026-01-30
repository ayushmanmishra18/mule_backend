from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/predict', methods=['POST'])
def predict():
    # This simulates your friend's deep neural network analysis
    data = request.json
    print(f"Received transaction {data.get('transaction_id')} for analysis...")
    
    # Mocking a risk score result
    return jsonify({
        "transaction_id": data.get("transaction_id"),
        "risk_score": 0.95,
        "recommendation": "ALLOW"
    })

if __name__ == '__main__':
    app.run(port=5000)