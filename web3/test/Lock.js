const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("MuleRegistry Contract", function () {
  let MuleRegistry;
  let muleRegistry;
  let owner;
  let bankA;
  let bankB;

  // Sample data for testing
  const txId = "TX-123456789";
  const senderAccount = "1234567890";
  const receiverAccount = "0987654321";
  const amount = ethers.parseEther("1.5"); // 1.5 ETH/Tokens
  const location = "New York, US";
  const reason = "Risk Score > 0.8";

  // Helper to hash accounts like the backend would
  const hashAccount = (account) => ethers.keccak256(ethers.toUtf8Bytes(account));

  beforeEach(async function () {
    // Get the ContractFactory and Signers here.
    MuleRegistry = await ethers.getContractFactory("MuleRegistry");
    [owner, bankA, bankB] = await ethers.getSigners();

    // Deploy a new contract before each test
    muleRegistry = await MuleRegistry.deploy();
    await muleRegistry.waitForDeployment();
  });

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      expect(await muleRegistry.owner()).to.equal(owner.address);
    });
  });

  describe("Reporting Mules", function () {
    it("Should allow a bank to report a mule transaction", async function () {
      const senderHash = hashAccount(senderAccount);
      const receiverHash = hashAccount(receiverAccount);

      // Bank A reports a mule
      await expect(
        muleRegistry.connect(bankA).reportMule(
          txId,
          senderHash,
          receiverHash,
          amount,
          location,
          reason
        )
      )
        .to.emit(muleRegistry, "MuleReported") // Check if event is emitted
        .withArgs(senderHash, receiverHash, amount, reason, (val) => val > 0); // Timestamp check ignored
    });

    it("Should flag both sender and receiver accounts", async function () {
      const senderHash = hashAccount(senderAccount);
      const receiverHash = hashAccount(receiverAccount);

      await muleRegistry.connect(bankA).reportMule(
        txId,
        senderHash,
        receiverHash,
        amount,
        location,
        reason
      );

      // Verify state update
      expect(await muleRegistry.isAccountFlagged(senderHash)).to.be.true;
      expect(await muleRegistry.isAccountFlagged(receiverHash)).to.be.true;
    });

    it("Should NOT flag unrelated accounts", async function () {
      const safeAccountHash = hashAccount("safe-account-id");
      expect(await muleRegistry.isAccountFlagged(safeAccountHash)).to.be.false;
    });
  });

  describe("Data Integrity", function () {
    it("Should store report details correctly", async function () {
      const senderHash = hashAccount(senderAccount);
      const receiverHash = hashAccount(receiverAccount);

      await muleRegistry.reportMule(txId, senderHash, receiverHash, amount, location, reason);

      // Fetch reports for sender
      const senderReports = await muleRegistry.accountReports(senderHash, 0);
      
      expect(senderReports.transactionId).to.equal(txId);
      expect(senderReports.reason).to.equal(reason);
      expect(senderReports.reporter).to.equal(owner.address);
    });

    it("Should increment report count for repeat offenders", async function () {
      const senderHash = hashAccount(senderAccount);
      const receiverHash = hashAccount(receiverAccount);

      // Report 1
      await muleRegistry.reportMule(txId, senderHash, receiverHash, amount, location, reason);
      
      // Report 2 (Different Tx, same mule)
      await muleRegistry.reportMule("TX-999", senderHash, receiverHash, amount, "London, UK", "Another Fraud");

      expect(await muleRegistry.getReportCount(senderHash)).to.equal(2);
    });
  });
});