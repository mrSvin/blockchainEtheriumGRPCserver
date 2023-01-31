// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

contract SimpleStorage {

    struct Wallet {
        string walletName;
        uint balance;
    }

    mapping(string => Wallet) public wallets;

    function setWallet(string memory nameWallet, uint setBalance) public {
        wallets[nameWallet].walletName = nameWallet;
        wallets[nameWallet].balance = setBalance;
    }

    function getWallet(string memory nameWallet) external view returns (Wallet memory) {
        return wallets[nameWallet];
    }

    function sendMoney(string memory nameWalletSender, string memory nameWalletRecipient, uint money) public {
        if (money <= wallets[nameWalletSender].balance && wallets[nameWalletRecipient].balance + money <= 4294967295) {
            wallets[nameWalletSender].walletName = nameWalletSender;
            wallets[nameWalletSender].balance = wallets[nameWalletSender].balance - money;

            wallets[nameWalletRecipient].walletName = nameWalletRecipient;
            wallets[nameWalletRecipient].balance = wallets[nameWalletRecipient].balance + money;
        }


    }

}