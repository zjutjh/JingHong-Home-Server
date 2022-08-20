import requests
import random
url = 'http://localhost:8888/api/new_normal'
for i in range(10):
    data = {
        'name': 'test' + str(i),
        'stu_id': str(random.randint(1, 100000000000)),
        'gender': random.choice((0, 1)),
        'college': 'college' + str(i),
        'campus': 'campus' + str(i),
        'phone': 'phone' + str(i),
        'qq': 'qq' + str(i),
        'region': random.choice((1, 2, 3)),
        'want1': random.choice((1, 2, 3, 4, 5, 6, 7, 8, 9)),
        'want2': random.choice((1, 2, 3, 4, 5, 6, 7, 8, 9)),
        'profile': 'profile' + str(i),
        'feedback': 'feedback' + str(i),
    }
    requests.post(url, data)
    print(data)
