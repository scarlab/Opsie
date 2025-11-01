import Config from '@/config'
import CsImage from '@/constants/image'
import { Button } from '../cn/button'
import { Bell } from 'lucide-react'
import NavUser from './NavUser'
import AddNew from './AddNew'
import { Link } from 'react-router-dom'
import TeamSwitcher from './TeamSwitcher'
import { Theme } from '../theme'
import { useViewContext } from '@/hooks/useViewContext'

export default function AppHeader() {
    const { title, subtitle, } = useViewContext();


    return (
        <nav className="bg-secondary border-b px-3 gap-3 flex items-center py-2 h-[var(--header-height)] sticky top-0 z-50">
            <Link to={'/'} className="flex items-center gap-1">
                <img className='drop-shadow w-8' width={400} height={400} src={CsImage.logo} alt={Config.projectName} />
                <h1 className='font-bold text-lg text-primary'> Opsie</h1>
            </Link>

            <div className="mx-2 h-7 w-0.5 bg-accent rotate-12" />

            <div className="grow flex gap-2 items-center">
                <TeamSwitcher />


                {title && <div className="mx-2 h-7 w-0.5 bg-accent rotate-12" />}

                {
                    title &&
                    <div className='ms-2'>
                        <h1 className="text-lg font-medium">{title}</h1>
                        {subtitle && <p className="text-xs text-muted-foreground/80">{subtitle}</p>}
                    </div>
                }
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
