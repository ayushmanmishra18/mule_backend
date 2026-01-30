// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MuleRegistry {
    // Event to notify listeners (other banks) of a new mule report
    event MuleReported(
        bytes32 indexed senderHash,
        bytes32 indexed receiverHash,
        uint256 amount,
        string reason,
        uint256 timestamp
    );

    // Structure to hold details of a specific fraud report
    struct FraudReport {
        string transactionId; // The original bank's internal ID
        uint256 amount;
        string location;
        uint256 timestamp;
        string reason;      // e.g., "High Risk Score > 0.8"
        address reporter;   // The wallet address of the bank reporting this
    }

    // Mapping to check if a specific account hash has been flagged
    // accountHash => bool isSuspicious
    mapping(bytes32 => bool) public flaggedAccounts;

    // Mapping to retrieve history of reports for a specific account hash
    mapping(bytes32 => FraudReport[]) public accountReports;

    address public owner;

    constructor() {
        owner = msg.sender;
    }

    /**
     * @dev Report a detected mule transaction.
     * @param _transactionId The unique ID from your backend (tx.TransactionID)
     * @param _senderHash SHA-256 hash of the sender's account number
     * @param _receiverHash SHA-256 hash of the receiver's account number
     * @param _amount The transaction amount
     * @param _location The location string from the transaction
     * @param _reason Context for the flag (e.g., "ML Score 0.95")
     */
    function reportMule(
        string memory _transactionId,
        bytes32 _senderHash,
        bytes32 _receiverHash,
        uint256 _amount,
        string memory _location,
        string memory _reason
    ) public {
        // Create the report struct
        FraudReport memory newReport = FraudReport({
            transactionId: _transactionId,
            amount: _amount,
            location: _location,
            timestamp: block.timestamp,
            reason: _reason,
            reporter: msg.sender
        });

        // Store the report for both sender and receiver
        accountReports[_senderHash].push(newReport);
        accountReports[_receiverHash].push(newReport);

        // Mark accounts as flagged for O(1) lookup
        flaggedAccounts[_senderHash] = true;
        flaggedAccounts[_receiverHash] = true;

        emit MuleReported(_senderHash, _receiverHash, _amount, _reason, block.timestamp);
    }

    /**
     * @dev Check if an account hash is in the blacklist.
     */
    function isAccountFlagged(bytes32 _accountHash) public view returns (bool) {
        return flaggedAccounts[_accountHash];
    }

    /**
     * @dev Get total reports for a specific account hash.
     */
    function getReportCount(bytes32 _accountHash) public view returns (uint256) {
        return accountReports[_accountHash].length;
    }
}