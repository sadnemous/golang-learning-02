from flask import Flask, request, jsonify
from flasgger import Swagger

app = Flask(__name__)
swagger = Swagger(app)

@app.route('/users', methods=['POST'])
def create_user():
    """
    Create a user.

    Accepts JSON only and returns JSON only.

    ---
    consumes:
      - application/json
    parameters:
      - name: body
        in: body
        required: true
        schema:
          type: object
          required:
            - name
            - email
          properties:
            name:
              type: string
            email:
              type: string
    responses:
      201:
        description: User created
        schema:
          type: object
          properties:
            id:
              type: integer
            name:
              type: string
            email:
              type: string
      400:
        description: Bad request (validation failed)
      415:
        description: Unsupported Media Type (Content-Type must be application/json)
    """
    if not request.is_json:
        return jsonify({'error': 'Content-Type must be application/json'}), 415

    body = request.get_json(silent=True)
    if not body:
        return jsonify({'error': 'Invalid JSON body'}), 400

    name = body.get('name')
    email = body.get('email')
    if not name or not email:
        return jsonify({'error': 'Both name and email are required'}), 400

    # For demo purposes we generate a simple id (in a real app use DB)
    import random
    user_id = random.randint(1000, 9999)

    user = {'id': user_id, 'name': name, 'email': email}
    return jsonify(user), 201

if __name__ == '__main__':
    app.run(debug=True)
