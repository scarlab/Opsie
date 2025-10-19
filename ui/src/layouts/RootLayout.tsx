import RootLoader from '@/components/loader/RootLoader';
import { Actions, useCsDispatch, useCsSelector } from '@/cs-redux';
import AuthSlice from '@/cs-redux/slices/auth.slice';
import { getLocalAuthUser, removeLocalAuthUser } from '@/helpers/auth.helper';
import { useEffect, useState } from 'react';
import { Outlet, useLocation, useNavigate } from 'react-router-dom';

export default function RootLayout() {
    const dispatch = useCsDispatch();
    const navigate = useNavigate();
    const location = useLocation();

    const { authUser, loading } = useCsSelector((state) => state.auth);

    const [sessionInitializing, setSessionInitializing] = useState(true);

    // Check session or restore from cache
    useEffect(() => {
        const init = async () => {
            const cachedUser = getLocalAuthUser();

            if (cachedUser && cachedUser.id && cachedUser.is_active) {
                dispatch(AuthSlice.actions.restoreAuthUser(cachedUser));
            } else {
                removeLocalAuthUser();
            }

            // Always verify server-side session (may take time)
            await dispatch(Actions.auth.session());

            setSessionInitializing(false);
        };


        const cachedUser = getLocalAuthUser();

        if (cachedUser && cachedUser.id && cachedUser.is_active) {
            dispatch(AuthSlice.actions.restoreAuthUser(cachedUser));
        } else {
            removeLocalAuthUser();
            dispatch(Actions.auth.session());
        }


        init();
    }, [dispatch]);

    // Redirect logic after session fully initialized
    useEffect(() => {
        if (!sessionInitializing && !loading) {
            if (!authUser && !location.pathname.startsWith('/auth')) {
                navigate('/auth/login', { replace: true });
            } else if (authUser && location.pathname.startsWith('/auth')) {
                navigate('/', { replace: true });
            }
        }
    }, [sessionInitializing, loading, authUser, navigate, location.pathname]);


    return (
        <div>
            {(sessionInitializing || loading)
                ?
                <RootLoader />
                :
                <Outlet />
            }
        </div>
    );
}
