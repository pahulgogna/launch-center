import { useEffect, useState } from "react"
import { types } from '../../wailsjs/go/models'
import { GetItems } from "../../wailsjs/go/storage/Handler"
import ItemCard from "../components/ui/ItemCard"

function Dashboard() {

    const [items, setItems] = useState<types.Item[]>([])

    async function populateItems() {
        let itms = await GetItems()
        console.log(itms)
        setItems(itms)
    }

    useEffect(() => {
        populateItems()
    }, [])

    return (
        <div className="flex gap-3">
            {items.map((it) => {
                return <ItemCard item={it} key={it.name}/>
            })}
        </div>
    )
}

export default Dashboard