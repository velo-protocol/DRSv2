pragma solidity ^0.5.0;

library LL {
    struct List {
        uint8 llSize;
        mapping (address => address) next;
    }

    function init(List storage l) internal returns (List memory) {
        l.next[address(1)] = address(1);

        return l;
    }

    function has(List storage l, address addr) internal view returns (bool) {
        return l.next[addr] != address(0);
    }

    function add(List storage l, address addr) internal returns (List memory) {
        require(!has(l, addr), "addr is already in the list");
        l.next[addr] = l.next[address(1)];
        l.next[address(1)] = addr;
        l.llSize++;

        return l;
    }

    function remove(List storage l, address addr, address prevAddr) internal returns (List memory) {
        require(has(l, addr), "addr not whitelisted yet");
        require(l.next[prevAddr] == addr, "wrong prevConsumer");
        l.next[prevAddr] = l.next[addr];
        l.next[addr] = address(0);
        l.llSize--;

        return l;
    }

    function getAll(List storage l) internal view returns (address[] memory) {
        address[] memory addrs = new address[](l.llSize);
        address curr = l.next[address(1)];
        for(uint256 i = 0; curr != address(1); i++) {
            addrs[i] = curr;
            curr = l.next[curr];
        }
        return addrs;
    }

    function getNextOf(List storage l, address curr) internal view returns (address) {
        return l.next[curr];
    }

}
