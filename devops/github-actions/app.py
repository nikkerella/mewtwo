from flask import Flask
import unittest

app = Flask(__name__)

@app.route('/')
def hello():
    return "Hello, World!"

def divide(a, b):
    return a / b

def safeDivide(a, b):
    try:
        return a / b
    except ZeroDivisionError:
        return "Error: Cannot divide by zero"

class TestApp(unittest.TestCase):
    def test_divide(self):
        self.assertEqual(divide(10, 2), 5)  # Normal case
        # with self.assertRaises(ZeroDivisionError):  # Test dividing by zero
        #     divide(1, 0)

        # divide(1, 0)
        safeDivide(1, 0)

if __name__ == '__main__':
    app.run(port=8080)

    # unittest.main()