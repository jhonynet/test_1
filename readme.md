// main blocks
Store: this persisit the entities, manages some state machine to allow modifications, and dispatch events when document
effectively change.

Indexer: takes incoming asset change events and index or re-index again those documents.


// entities
type Asset {
ID int64

    Name string
    ticker string
    
    supply int
    price float64

    meta: map[string]interface{} {
        // withdraw-related attributes
        "withdraw.enabled": true,
        "withdraw.min_deposit_amount": 100,
        "withdraw.max_deposit_amount": 100,
        // deposit-related attributes
        "deposit.enabled": true,
        "deposit.min_deposit_amount": 100,
        // real world assets
        "stock.enabled": true,
        // display rules
        'visibility': true
    }

}

// data path
1 - store it in Postgres as JSONB
2 - extract indexable fields
3 - index the ID of the document, along with its indexable fields.

// search query
1 - Search -> handled by indexer -> return IDs

2 - Get By IDs -> Lookup O1

