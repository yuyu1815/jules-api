# Jules API Test Suite

This directory contains a comprehensive test suite for the Jules API. It is used to verify that the client libraries in each language are functioning correctly.

## 📁 File Structure

- `test_api.py` - Test program for the Python client
- `test_api_js.js` - Test program for the JavaScript/TypeScript client
- `test_api_go.go` - Test program for the Go client
- `.env` - Environment variable file containing the API key
- `README.md` - Description of this test suite

## 🔐 Security Configuration

### Environment Variable File (.env)

To run the tests, you need to set the API key in the environment variable file:

```bash
# test/.env
JULES_API_KEY=your_actual_api_key_here
```

**Important**: Do not commit this `.env` file to version control.

## 🚀 How to Run Tests

### Python Tests

```bash
# Prerequisite: Python environment and required packages are installed
cd py
pip install python-dotenv requests pydantic
cd ../test
python3 ../test/test_api.py
```

### JavaScript Tests

```bash
# Prerequisite: Node.js and npm are installed
# Not yet verified (environment dependent)
cd test
node test_api_js.js
```

### Go Tests

```bash
# Prerequisite: Go is installed
# Not yet verified (environment dependent)
cd go
go run ../test/test_api.go
```

## 🧪 Test Content

Each test program tests the following API endpoints:

1. **📋 List Sources** - Retrieve list of available sources (GitHub repositories)
2. **🚀 Create Session** - Create a new session
3. **📖 Get Session** - Get details of a specific session
4. **📂 List Sessions** - Retrieve list of sessions
5. **🎬 List Activities** - Retrieve list of activities in a session
6. **💬 Send Message** - Send a message to the agent
7. **📦 Get Source** - Get details of a specific source

## 📊 Expected Results

- **Normal cases**: Most tests (6/7) should succeed
- **Limitations**: Send Message test may fail due to session initialization wait
- **Verification content**: Confirm API authentication, HTTP communication, and data structure correctness

## 🎯 Test Purpose

1. **API Functionality Verification** - Confirm that each endpoint operates normally
2. **Client Implementation Validity** - Confirm that libraries in each language are correctly implemented
3. **Error Handling** - Confirm that appropriate error handling is performed
4. **Security** - Confirm that API keys are handled securely

## ⚠️ Notes

- API tests connect to the actual API server, so be mindful of rate limits
- Running tests will actually create Jules sessions
- If errors occur, check network connectivity and API key validity

## 📞 Support

For issues with the tests, check the Jules API documentation or contact the support team.
