pragma solidity ^0.5.0;

contract MockIERC20 {
    function totalSupply() external view returns (uint256){
        return 0;
    }

    function balanceOf(address) external view returns (uint256){
        return 0;
    }

    function transfer(address, uint256) external returns (bool){
        return true;
    }

    function allowance(address, address) external view returns (uint256){
        return 0;
    }

    function approve(address, uint256) external returns (bool){
        return true;
    }

    function transferFrom(address, address, uint256) external returns (bool){
        return true;
    }

}
