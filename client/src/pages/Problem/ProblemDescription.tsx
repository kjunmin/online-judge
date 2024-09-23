import { useQuery } from "@tanstack/react-query";
import axios from "axios";

type Problem = {
    problemID: string;
    title: string;
    description: string; 
    tags: string[];
}

function useProblem(problemId: string) {
    return useQuery({
        queryKey: [problemId],
        queryFn: async () => await axios.get<never, Problem>(`/api/v1/problem/${problemId}`)
    })
}

type ProblemDescriptionProps = {
    problemId: string
}

export function ProblemDescription({ problemId }: ProblemDescriptionProps) {
    const { data } = useProblem(problemId)

    return (
        <div>
            {JSON.stringify(data)}
        </div>
    )
}