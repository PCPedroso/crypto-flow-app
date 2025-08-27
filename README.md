# Cryptocurrency Transaction Application

This is a cryptocurrency transaction application built using Go. The application allows users to process cryptocurrency transactions and check the health of the application.

## Project Structure
```
crypto-transaction-app
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── handlers.go  # HTTP request handlers
│   ├── models
│   │   └── models.go    # Data structures for transactions and users
│   ├── routes
│   │   └── routes.go    # Application routes
│   └── services
│       └── services.go  # Business logic for transactions
├── go.mod                # Module definition file
├── go.sum                # Dependency checksums
└── README.md             # Project documentation
```

## Features

*   **Cryptocurrency Transactions:** Process cryptocurrency transactions securely and efficiently.
*   **Multiple Wallet Management:** Support for integrating various wallets, allowing centralized asset management.
*   **Reports and History:** Generate detailed transaction reports for accounting or auditing purposes, with export options in practical formats.
*   **Two-Factor Authentication (2FA):** Enhance the security of your transactions with additional layers of protection.
*   **Trusted Wallet Registry:** A registry of wallets that can be used in transactions without requiring 2FA authentication, as long as they are within the limit defined in the registry.
*   **Real-Time Conversion:** View the value of transactions in fiat currencies, with updated quotes for easy financial management.
*   **Customizable Notifications:** Receive instant alerts about receipts, payments, or suspicious activity.
*   **Health Check:** Check the health status of the application.

## Setup Instructions

1.  **Clone the repository:**
    ```
    git clone https://github.com/PCPedroso/crypto-flow-app.git
    cd crypto-transaction-app
    ```

2.  **Install dependencies:**
    ```
    go mod tidy
    ```

3.  **Run the application:**
    ```
    go run cmd/main.go
    ```

## Usage Guidelines

*   The application exposes HTTP endpoints for processing transactions, managing wallets, and checking the health status.
*   You can use tools like Postman or curl to interact with the API.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.