pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
import "./VeloToken.sol";
import "./WhitelistAssets.sol";

contract Warp is WhitelistAdminRole{

    WhitelistAssets private whitelistAssets;

    uint16 private _threshold;

    constructor(address whitelistAssetsAddr, uint16 threshold) public {
        whitelistAssets = WhitelistAssets(whitelistAssetsAddr);
        _threshold = threshold;
    }

    event Unlock(address indexed to, bytes32 indexed assetName, uint256 tokens);
    event Lock(address indexed tokenOwner, bytes32 indexed assetName, uint256 tokens);

    /**
    *   unlock transfers some valid token to the destination address. The same amount of token on the stellar
    *   wiil be locked in the stellar unlockup account.
    **/
    function unlock(address to, bytes32 assetName, uint256 tokens) public onlyWhitelistAdmin {

        VeloToken(whitelistAssets.creditAddr(assetName)).mint(to, tokens);

        emit Unlock(to, assetName, tokens);

    }

    /**
    *   lock burns a tokens on the EvryNet chain and pay the same amount of token in the Lockup account to the
    *   token owner account on Stellar.
    **/
    function lock(bytes32 assetName, uint256 tokens) public {

        VeloToken(whitelistAssets.creditAddr(assetName)).burn(msg.sender, tokens);

        emit Lock(msg.sender, assetName, tokens);

    }

}
