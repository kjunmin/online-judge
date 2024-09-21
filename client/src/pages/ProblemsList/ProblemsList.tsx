import styles from './ProblemsList.module.css'
import { Problem } from "../types";
import axios from "axios";
import { useQuery } from "@tanstack/react-query";

type ProblemListItem = Pick<Problem, 'problemID' | 'tags' | 'title'>

function useGetProblemsList() {
    return useQuery<ProblemListItem[]>({
        queryKey: ['getProblemsList'],
        queryFn: async () => await axios.get('/api/v1/problemslist')
    })
}

export function ProblemsList() {

    const { isLoading, data: problemsList } = useGetProblemsList();

    if (isLoading) {
        return <h2>Loading...</h2>
    }

    return (
        <div className={styles.container}>
            <ul>
                {problemsList?.map(problemsListItem => {
                    return (
                        <li>
                            <p>{problemsListItem.problemID}</p>
                            <p>{problemsListItem.title}</p>
                            {problemsListItem.tags.map(tag => {
                                return <p>{tag}</p>
                            })}
                        </li>
                    )
                })}
            </ul>
        </div>
    )
}