package main

import (
	"os"

	"errors"

	"strings"

	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/fapiko/route53-dynamic-dns/ip"
)

func main() {
	config, err := parseConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(config.AwsAccessKeyID, config.AwsSecretAccessKey, ""),
	})
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	svc := route53.New(sess)

	for true {
		externalAddr, err := ip.External()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		err = upsertARecord(svc, config.Hostname, externalAddr)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		time.Sleep(300 * time.Second)
	}
}

func upsertARecord(svc *route53.Route53, name string, value string) error {
	hostname := strings.SplitN(name, ".", 2)

	zoneID, err := getHostedZoneID(svc, hostname[1])
	if err != nil {
		return err
	}

	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(name),
						Type: aws.String("A"),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(value),
							},
						},
						TTL: aws.Int64(300),
					},
				},
			},
		},
		HostedZoneId: aws.String(zoneID),
	}

	_, err = svc.ChangeResourceRecordSets(params)
	if err != nil {
		return err
	}

	return nil
}

func getHostedZoneID(svc *route53.Route53, name string) (string, error) {
	if !strings.HasSuffix("name", ".") {
		name = name + "."
	}

	hostedZones, err := svc.ListHostedZones(&route53.ListHostedZonesInput{})
	if err != nil {
		return "", err
	}

	for _, zone := range hostedZones.HostedZones {
		if strings.ToLower(*zone.Name) == strings.ToLower(name) {
			return *zone.Id, nil
		}
	}

	return "", errors.New("zone not found for " + name)
}
