import Config from '@/config'
import CsImage from '@/constants/image'
import { Button } from '../cn/button'
import { Bell, HelpingHand } from 'lucide-react'
import NavUser from './NavUser'
import AddNew from './AddNew'

export default function AppHeader() {
    return (
        <nav className="bg-secondary border-b px-3 gap-3 flex items-center py-2 h-[var(--header-height)] sticky top-0 z-50">
            <div className="w-9 aspect-square rounded-full overflow-hidden ">
                <img width={200} height={200} src={CsImage.logo} alt={Config.projectName} />
            </div>

            <div className="grow flex gap-2 items-center">
                <h3 className='text-lg font-black text-primary'>Opsie</h3>
            </div>
            <div className='flex gap-2 items-center'>
                <AddNew />
                <Button variant={'outline'} size={'icon'}><HelpingHand /></Button>
                <Button variant={'outline'} size={'icon'}><Bell /></Button>
                <NavUser />
            </div>
        </nav>
    )
}
