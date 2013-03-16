These are just rough drafts and will most likely expand/change in the future.




When it comes to implementing intent handlers for languages that I have no interest in myself, as I see it people are free to choose how they want to do it. If someone wants to merge code into the main repository for new features or languages, I'll require that if it's merged that they also commit to maintaining that code and fixing bugs in it as they are revealed. And with that perspective maybe the path of least resistance for those authors is to reuse existing resources written in another language than Go or maybe it is to write new code from scratch. If we figure out how to implement this "register myself as an intent handler" mechanism maybe Go could wouldn't have to be touched at all.

Ideally code to be merged into the main repository would be written in pure Go code without reliance on external dependencies to make sure that the full regression test suite runs everywhere, across all platforms, no matter what particular versions of other software/libraries are installed on a particular machine if they are installed at all.

I understand that that isn't always possible or would be a huge effort, so in the instances where external dependencies are used, I'll require that the full test suite should run on travis-ci (#31) as well as if the required external dependencies are not installed, the code should not register the intent handlers associated nor should those specific tests run. Currently that is detected at runtime, but could be handled via a build tag too.
