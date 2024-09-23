import { useParams } from "react-router-dom"
import { isNil } from 'lodash'
import { ProblemDescription } from "./ProblemDescription"

type ProblemParams = {
    problemId: string
}

export function Problem() {
    const { problemId } = useParams<ProblemParams>()

    if (isNil(problemId)) {
        return <>Error empty ID</>
    }

    return (
        <div>
            <h1>Problem Page</h1>
            <ProblemDescription problemId={problemId} />
        </div>
    )
}