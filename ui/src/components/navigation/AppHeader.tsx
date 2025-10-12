import Config from '@/config'
import CsImage from '@/constants/image'
import { Button } from '../cn/button'
import { Bell } from 'lucide-react'
import NavUser from './NavUser'
import AddNew from './AddNew'
import { Link } from 'react-router-dom'
import OrganizationSwitcher from './OrganizationSwitcher'
import { Theme } from '../theme'

export default function AppHeader() {
    return (
        <nav className="bg-secondary border-b px-3 gap-3 flex items-center py-2 h-[var(--header-height)] sticky top-0 z-50">
            <Link to={'/'} className="w-8 aspect-square rounded-full overflow-hidden ">
                <img width={200} height={200} src={CsImage.logo} alt={Config.projectName} />
            </Link>

            <div className="mx-2 h-7 w-0.5 bg-accent rotate-12"></div>

            <div className="grow flex gap-2 items-center">
                <OrganizationSwitcher />
            </div>

            <div className='flex gap-3 items-center'>
                {Config.isDev && <Theme />}
                <div className='flex items-center gap-2 border rounded-2xl px-2.5 py-0.5 bg-accent/60 cursor-pointer'>
                    <div className='w-2 rounded-full aspect-square bg-green-500' />
                    <span className='text-xs'>All Ok</span>
                </div>
                <Button variant={'outline'} size={'icon'}><Bell /></Button>
                <AddNew />
                <NavUser />
            </div>
        </nav>
    )
}
