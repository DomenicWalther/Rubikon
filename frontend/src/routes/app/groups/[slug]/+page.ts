
import type { PageLoad } from './$types';

export const load: PageLoad = ({params}) => {
    return { group_id: params.slug }
}