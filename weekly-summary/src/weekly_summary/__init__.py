import git
from datetime import datetime, timedelta
import os

def get_weekly_diff():
    repo = git.Repo('..')

    last_week = datetime.now() - timedelta(days=7)
    commits = list(repo.iter_commits(since=last_week))

    if(len(commits) == 0):
        return Null
    
    diffs = commits[-1].diff(commits[0], None, True)

    parsed_diffs = []

    for diff in diffs:
        diffString = diff.diff.decode()

        if diffString.startswith("Binary files ") or len(diffString) == 0:
            continue

        parsed_diffs.append({
            "path": diff.a_path,
            "diff": diff.diff.decode(),
        })

    return parsed_diffs
