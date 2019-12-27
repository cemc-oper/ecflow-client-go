#pragma once

#include <string>
#include <vector>

namespace EcflowUtil {

struct NodeStatusRecord {
    std::string path_;
    std::string status_;
};

class EcflowClientWrapperPrivate;

class EcflowClientWrapper {
public:
    EcflowClientWrapper() = delete;
    EcflowClientWrapper(const std::string &host, const std::string &port);
    ~EcflowClientWrapper();

    void setConnectTimeout(int time_out);

    int sync();

    std::vector<NodeStatusRecord> statusRecords() {
        return status_records_;
    }

    std::string statusRecordsJson();
    const char* statusRecordsJsonChar();

    std::string errorMessage();

private:
    std::string host_;
    std::string port_;

    EcflowClientWrapperPrivate* p_;
    std::vector<NodeStatusRecord> status_records_;
};

}

