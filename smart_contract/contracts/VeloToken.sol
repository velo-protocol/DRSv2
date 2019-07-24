pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";
import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";

contract VeloToken is ERC20, ERC20Detailed, WhitelistAdminRole {

    // stellar address who issuing this credit
    string private _issuer;

    constructor (
        string memory issuer,
        string memory name,
        string memory code,
        uint8 decimals
    )
    public
    ERC20Detailed(name, code, decimals) {
        _issuer = issuer;
    }

    /**
    * Return the credit issuer address (stellar address)
    */
    function issuer() public view returns (string memory) {
        return _issuer;
    }

    function mint(address account, uint256 amount) public onlyWhitelistAdmin returns (bool) {
        _mint(account, amount);
        return true;
    }

    function burn(address account, uint256 tokens) public onlyWhitelistAdmin {
        _burn(account, tokens);
    }
}
