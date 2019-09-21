#pragma once

#include <string>
#include <vector>

namespace EcflowUtil {

struct NodeStatusRecord {
    std::string path_;
    std::string status_;
};

class EcflowClientPrivate;

class EcflowClient {
public:
    EcflowClient() = delete;

    EcflowClient(const std::string &host, const std::string &port);

    ~EcflowClient();

    int sync();

    std::vector<NodeStatusRecord> statusRecords() {
        return status_records_;
    }

    std::string errorMessage();

private:
    std::string host_;
    std::string port_;

    EcflowClientPrivate* p_;
    std::vector<NodeStatusRecord> status_records_;
};

}

