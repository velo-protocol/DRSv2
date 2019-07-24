pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";

contract WhitelistAssets is WhitelistAdminRole {
    
    mapping(bytes32=>address) public assets;

    event AddedAsset(bytes32 indexed assetName, address indexed assetAddr);
    event RemovedAsset(bytes32 indexed assetName);

    /**
    *   addAssets adds a new supported asset to the asset whitelist
    **/
    function addAsset(bytes32 assetName, address assetContract) public onlyWhitelistAdmin {

        assets[assetName] = assetContract;
        
        emit AddedAsset(assetName, assetContract);

    }

    function removeAsset(bytes32 assetName) public onlyWhitelistAdmin {

        assets[assetName] = address(0);

        emit RemovedAsset(assetName);

    }

    /**
    *   creditAddr returns the credit contract address which identified by asset name
    **/
    function creditAddr(bytes32 assetName) public view returns (address assetAddr) {

        require(assets[assetName] != address(0), "WhitelistAssets: the asset is prohibited");
        assetAddr = assets[assetName]; 

    }
}
