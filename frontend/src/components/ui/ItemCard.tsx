import { types } from '../../../wailsjs/go/models'
import { BrowserOpenURL } from "../../../wailsjs/runtime/runtime";

function handleRedirect(url: string) {
  BrowserOpenURL(url)
}

function ItemCard({ item }: { item: types.Item }) {
  return (
    <div className='w-fit bg-white/50 p-5 border rounded-md'>
      <div onClick={() => handleRedirect(item.url)} className='hover:cursor-pointer'>
        <div className='flex justify-center w-full mb-3'>
          <img className='rounded-lg' src={item.thumbnail} height={100} width={300}/>
        </div>
        <div className='flex font-semibold text-xl justify-center'>
          {item.name}
        </div>
        <div className='flex justify-center text-balance mt-2'>
          {item.description}
        </div>
      </div>
    </div>
  )
}

export default ItemCard