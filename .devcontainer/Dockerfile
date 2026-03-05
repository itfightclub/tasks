# ------------------------------------------------------------
# Base image – official Go (security‑maintained)
# ------------------------------------------------------------
FROM docker.io/library/golang:1.26.0

# ------------------------------------------------------------
# Args/Variables
# ------------------------------------------------------------
ARG USERNAME=tasks
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# ------------------------------------------------------------
# Install any & all packages in one action 
# ------------------------------------------------------------
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        bash less man nano sudo towncrier tree \
    && rm -rf /var/lib/apt/lists/* 

# Reasoning:
    # bash = interactive dev work
    # less = interactive dev work
    # man = manual pages
    # nano = git commit editor
    # sudo = changes while developing
    # towncrier = changelog tool
    # tree = documentation & dev work

# ------------------------------------------------------------
# Give the $USERNAME password‑less sudo. 
# Use bash over dash (default) for interactive features:
# History (up arrow), auto-complete (tab), color, etc
# ------------------------------------------------------------
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m -s /bin/bash $USERNAME \
    && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# ------------------------------------------------------------
# Working directory for the source code (will be bind‑mounted)
# ------------------------------------------------------------
WORKDIR /workspace

# ------------------------------------------------------------
# Declare that the container expects a volume at the user home.
# This is only documentation – the actual named volume is attached
# via devcontainer.json.  Declaring it here prevents Docker from
# creating an anonymous volume if the container is started without
# the explicit -v flag.
# ------------------------------------------------------------
VOLUME ["/home/$USERNAME"]

# ------------------------------------------------------------
# Switch to the non‑root user for everything that runs later
# ------------------------------------------------------------
USER $USERNAME

# Keep the container alive so VS Code can attach
CMD ["sleep", "infinity"]
