import random
import string
import time


def app(environ, start_response): 
    print("Received request:"
          f"\nMethod: {environ['REQUEST_METHOD']}"
          f"\nPath: {environ['PATH_INFO']}"
          f"\nUser Agent: {environ['HTTP_USER_AGENT']}")

    length = random.randint(8, 16)
    required_symbols = "#.,!@&^%*"

    password_chars = [
        random.choice(string.digits),
        random.choice(string.ascii_lowercase),
        random.choice(string.ascii_uppercase),
        random.choice(required_symbols),
    ]

    all_allowed_chars = string.digits + string.ascii_lowercase + string.ascii_uppercase + required_symbols
    password_chars.extend(random.choice(all_allowed_chars) for _ in range(length - len(password_chars)))
    random.shuffle(password_chars)

    password = "".join(password_chars).encode("utf-8")
    time.sleep(0.05)

    start_response("200 OK", [
        ("Content-Type", "text/plain"),
        ("Content-Length", str(len(password)))
    ])
    return iter([password])