const { network  ,getNamedAccounts, deployments } = require("hardhat");
const { networkConfig, developmentChains } = require("../helper-hardhar-config");

module.exports = async () => { 
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();
  const chainId = network.config.chainId;

  log("\n=============== Starting Deployment ===============");
  log(`Network: ${network.name} (Chain ID: ${chainId})`);
  log(`Deployer: ${deployer}`);

  try {
    const uploadContract = await deploy("MuleRegistry", {
      from: deployer,
      args: [], 
      log: true,
      waitConfirmations: networkConfig[chainId]?.blockConfirmations || 1,
    });

    log(`✅ Contract deployed to: ${uploadContract.address}`);
    log("=================================================\n");
  } catch (error) {
    console.error("\n❌ Deployment Failed!");
    console.error(error);
    process.exit(1);
  }
};

// Add tags to target this script for execution
module.exports.tags = ["upload"];