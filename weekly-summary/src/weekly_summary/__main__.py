import streamlit as st
from get_weekly_diff import get_weekly_diff

st.title('Weekly Summary')

weekly_diffs = get_weekly_diff()

if weekly_diffs is None:
    st.write("No changes in the last week")

for diff in weekly_diffs:
    st.subheader(diff['path'])
    st.code(diff['diff'], language='diff')
