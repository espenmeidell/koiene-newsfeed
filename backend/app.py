from chalice import Chalice
import uuid
import time

app = Chalice(app_name='backend')


@app.route('/')
def index():
    return {'hello': 'world'}


POSTS = [
    {
        'id': '6d503941-8d24-437c-9f19-88c0c6dce48b',
        'title': 'Complaints at Fosenkoia',
        'content': "We have received some complaints from neighbours at Fosen. Don't walk through private gardens! When walking from the speed boat you need to go up to the main road behind the store. The road along the docks is a dead end and will lead you into private gardens.",
        'status': 'ACTIVE',
        'timestamp': int(time.time())
    },
    {
        'id': '5e7995cb-7422-4b2d-a5ed-8907a3bbe435',
        'title': 'Summer Reservation Period',
        'content': "The summer reservation period starts may 30th. From that date on you can reserve cabins between june 8th and 23rd august. Akademika is not open all summer, so remember to check their opening hours so you can pick up keys.",
        'status': 'ACTIVE',
        'timestamp': int(time.time())
    }
]


def get_single_post(id):
    return list(filter(lambda p: p['id'] == id, get_all_posts()))[0]


def get_all_posts():
    return POSTS


def persist_post(post):
    POSTS.append(post)


@app.route('/posts')
def get_posts():
    return get_all_posts()


@app.route('/posts/{id}')
def get_post(id):
    return get_single_post(id)


@app.route('/posts', methods=['POST'])
def add_post():
    post = app.current_request.json_body
    post['status'] = "ACTIVE"
    post['id'] = str(uuid.uuid4())
    post['timestamp'] = int(time.time())
    persist_post(post)
    return post

# The view function above will return {"hello": "world"}
# whenever you make an HTTP GET request to '/'.
#
# Here are a few more examples:
#
# @app.route('/hello/{name}')
# def hello_name(name):
#    # '/hello/james' -> {"hello": "james"}
#    return {'hello': name}
#
# @app.route('/users', methods=['POST'])
# def create_user():
#     # This is the JSON body the user sent in their POST request.
#     user_as_json = app.current_request.json_body
#     # We'll echo the json body back to the user in a 'user' key.
#     return {'user': user_as_json}
#
# See the README documentation for more examples.
#
