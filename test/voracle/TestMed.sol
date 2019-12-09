//pragma solidity ^0.5.0;
//
//import "../../contracts/voracle/Med.sol";
//import "../../contracts/price-feeds/FeedsFactory.sol";
//import "../../contracts/book-room/LL.sol";
//
//contract TestMed {
//    FeedsFactory public feedsFactory;
//    Med public med;
//
//    LL.List public feeders;
//    using LL for LL.List;
//
//    constructor() public {
//        feedsFactory = new FeedsFactory();
//
//        feeders = feeders.init();
//        feeders.add(feedsFactory.create());
//        feeders.add(feedsFactory.create());
//
//        med = new Med(address(this), "XLMUSD");
//    }
//}
