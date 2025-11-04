import Role from "@/components/utils/Role";
import { Actions, useCsDispatch, useCsSelector } from "@/cs-redux";
import { Check, X } from "lucide-react";
import { useEffect } from "react";
import { Link } from "react-router-dom";

export default function UserList() {
    const dispatch = useCsDispatch();
    const { users } = useCsSelector(state => state.user);


    useEffect(() => {
        dispatch(Actions.user.GetAllUser());
    }, [])


    return (
        <div>
            <div className="mt-3">
                <div className="grid grid-cols-12 gap-2 px-3 pb-1 text-muted-foreground text-sm border-b">
                    <p className="col-span-1"></p>
                    <p className="col-span-3">Name</p>
                    <p className="col-span-5">Email</p>
                    <p className="col-span-2 text-center">System Role</p>
                    <p className="col-span-1 text-right">Status</p>
                </div>
                {users ?
                    users.map((user, i) => (
                        <Link to={`/settings/user?user=${user.id}`} key={i} className="grid grid-cols-12 items-center hover:bg-accent/20 transition-all py-1.5 border-b px-3">

                            <div className="col-span-1">
                                <div className="w-9 aspect-square rounded-full grid place-items-center" >
                                    <img width={100} height={100} src={user.avatar} />
                                </div>
                            </div>
                            <p className="col-span-3 font-medium">{user.display_name}</p>
                            <p className="col-span-5 text-muted-foreground text-sm">{user.email}</p>
                            <div className="col-span-2 text-xs flex justify-center ">
                                <Role sysRole={user.system_role} />
                            </div>
                            <div className="col-span-1 w-full flex justify-end items-center">
                                {user?.is_active ? <Check size={19} className="text-green-500" /> : <X size={19} color='red' />}
                            </div>
                        </Link>
                    ))
                    :
                    <div>
                        Loading...
                    </div>
                }
            </div>
        </div>
    )
}
