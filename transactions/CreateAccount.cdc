// PUBLIC KEY 1 =>
// 9a2850585a0ef9741d2c2109a9fcf43a49d0e9f5cccd9c17494762a9070a183b5833b21c50ab7126f4992db93ac0f234f7ae15e339280c21d83de2112f73bf47

// PUBLIC KEY 2 =>
// 01fc65e9febf9b4c3bc67a4cf7bbf17a4d441cd7af6c6a6a7a72fec37433be48a0a4e72bc048e76dbf42a7fc707ee50399ba7fcd1028abd9436ee8a7ad8f0d78
transaction(publicKey: String) {

    prepare(signer: AuthAccount) {

        let key = PublicKey(
            publicKey: publicKey.decodeHex(),
            signatureAlgorithm: SignatureAlgorithm.ECDSA_P256 
        )

        let account = AuthAccount(payer: signer)

        account.keys.add(
            publicKey: key,
            hashAlgorithm: HashAlgorithm.SHA3_256,
            weight: 1000.0
        )
    }

    execute {
        log("successfully created an account")
    }

}