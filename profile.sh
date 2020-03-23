#!/usr/bin/env bash

# Just gets the top level directory of this project. Useful for scripting within the project via relative file paths
KUBERT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

kubert () {
    # if no command given force help page
    local OPTION
	if [[ "$1" != "" ]]; then
        OPTION=$1
    else
        OPTION="help"
    fi
	# handle input options
    case "${OPTION}" in
        'help')
echo "Usage: $ ${FUNCNAME} [option] [flags]
Options:
- help: show this menu
- mock: Mock all services (req: gomock)
- test: run all mock tests
"
        ;;
        'mock')
            kubertMockServices
        ;;
        'test')
          kubertTestServers
        ;;
        *)
            echo -e "ERROR: invalid option. Try..\n$ ${FUNCNAME} help"
        ;;
    esac
}

kubertTestServers () {
     go test $(go list $KUBERT_DIR/...)
}

# Generate mock files for all services, putting the results in the proper file. renames some stuff for consistency.
# If you update any services, recommend running this function to update the services for the tests.
kubertMockServices () {
    packageName="mockgen"

    MOCK_FOLDER="services"
    SERVICE_DIR="${KUBERT_DIR}/${MOCK_FOLDER}"
    SERVICES=$(find "${SERVICE_DIR}" -maxdepth 1 -mindepth 1 -type d)
    for SERVICE_PATH in ${SERVICES}
    do
        if [[ -f ${SERVICE_PATH}/interface.go ]]; then
            FOLDER_NAME="${SERVICE_PATH##*/}"
            mockgen \
                -source=${SERVICE_PATH}/interface.go \
                -destination=mocks/${MOCK_FOLDER}_mocks/${FOLDER_NAME}_mock.go \
                -package=${MOCK_FOLDER}_mocks \
                -mock_names Service=Mock_${FOLDER_NAME}
        else
          PROTOS=$(find "${SERVICE_PATH}" | grep ".pb.go")
          for PROTO in ${PROTOS}
          do
            PROTO_FILE_NAME=$(basename "${PROTOS}")
            PROTO_FILE_NAME_STRIP_EXT=${PROTO_FILE_NAME/.go/}
            PROTO_REPLACED_NAME=${PROTO_FILE_NAME_STRIP_EXT/./_}
            mockgen \
                -source=${SERVICE_PATH}/${PROTO_FILE_NAME} \
                -destination=mocks/${MOCK_FOLDER}_mocks/"${PROTO_REPLACED_NAME}"_mock.go \
                -package=${MOCK_FOLDER}_mocks \
                -mock_names Service=Mock_"${PROTO_REPLACED_NAME}"
          done
        fi
    done
}

# Check if a command exists in the environment
# Returns 0 if command found
package-installed () {
	result=$(compgen -A function -abck | grep "^$1$")
    # Note that in bash, non-zero exit codes are error codes. returning 0 means success
	if [[ "${result}" == "$1" ]]; then
		# package installed
		return 0
	else
		# package not installed
		return 1
	fi
}
